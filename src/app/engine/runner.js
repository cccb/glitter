
import {to_luastring,

        lua,
        luaconf,
        lauxlib,
        lualib} from "fengari";

import Framebuffer from './framebuffer'

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
  }
  
  /*
   * Load script and execute
   */
  loadScript(script) {
    const luaScript = to_luastring(script);
    if (lauxlib.luaL_dostring(this.L, luaScript) != 0) {
      return lua.lua_tostring(this.L, -1) // Err
    }
  }

  /*
   * Call render function provided by lua script
   */
  render() {
    if(!lua.lua_getglobal(this.L, "render")) {
      throw "render() function missing in script";
      return;
    }
    lua.lua_pushlightuserdata(this.L, this.framebuffer);
    lauxlib.luaL_setmetatable(this.L, "framebuffer");
    lua.lua_pushnumber(this.L, 23.0);

    lua.lua_pcall(this.L, 2, 0, 0);
  }

}


