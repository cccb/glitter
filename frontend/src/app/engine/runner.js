
import {to_luastring,
        to_jsstring,

        lua,
        luaconf,
        lauxlib,
        lualib} from "fengari";

import Framebuffer from './framebuffer'
import LuaBootScript from './lua-boot'

export default class Runner {
  /*
   * Initialize runner:
   *  - Create Lua state and setup framebuffer
   */
  constructor() {
    // Create Lua state
    this.L = lauxlib.luaL_newstate();
    lualib.luaL_openlibs(this.L);

    // Initialize buffer
    this.framebuffer = new Framebuffer(this, 1, 14);

    // Load bootstrapping
    this.loadScript(LuaBootScript);
  }

  /*
   * Load script and execute
   */
  loadScript(script) {
    const luaScript = to_luastring(script);
    if (lauxlib.luaL_dostring(this.L, luaScript) != 0) {
      let err = lua.lua_tostring(this.L, -1);
      console.error(to_jsstring(err));
      return err;
    }
  }

  /*
   * Call render function provided by lua script
   */
  render(t) {
    if(!lua.lua_getglobal(this.L, "render")) {
      throw "render() function missing in script";
      return;
    }
    lua.lua_pushlightuserdata(this.L, this.framebuffer);
    lauxlib.luaL_setmetatable(this.L, "framebuffer");
    lua.lua_pushnumber(this.L, t);
    lua.lua_pcall(this.L, 2, 0, 0);
  }

}


