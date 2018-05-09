
import {to_luastring,

        lua,
        luaconf,
        lauxlib,
        lualib} from "fengari";

export default class Framebuffer {

  constructor(engine, width, height) {
    this.width = width;
    this.height = height;

    this.buffer = new Array(width * height);
  
    // Register framebuffer library
    if (lauxlib.luaL_newmetatable(engine.L, "framebuffer") != 0) {
      lua.lua_pushjsfunction(engine.L, function(L){
        console.log("FNOOORF");
      });
      lua.lua_setfield(engine.L, -2, "debug");

      lua.lua_pushvalue(engine.L, -1);
      lua.lua_setfield(engine.L, -2, "__index");
    }
  }

  luaDebug(L) {
    console.log("luaDebug");
    console.log(fb);
    const fb = lua.lua_touserdata(L);
  }

}


