
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
      lua.lua_pushjsfunction(engine.L, (L) => this.luaDebug(L));
      lua.lua_setfield(engine.L, -2, "debug");

      lua.lua_pushjsfunction(engine.L, (L) => this.luaGetpixel(L));
      lua.lua_setfield(engine.L, -2, "getpixel");

      lua.lua_pushjsfunction(engine.L, (L) => this.luaSetpixel(L));
      lua.lua_setfield(engine.L, -2, "setpixel");

      lua.lua_pushjsfunction(engine.L, (L) => this.luaGetsize(L));
      lua.lua_setfield(engine.L, -2, "getsize");

      lua.lua_pushvalue(engine.L, -1);
      lua.lua_setfield(engine.L, -2, "__index");
    }
  }

  getPixel(x, y) {
    return this.buffer[x % this.width + y * this.height];
  }

  setPixel(x, y, v) {
    if (x < 0 || x > this.width || y < 0 || y > this.height) {
      return null;
    }

    this.buffer[x % this.width + y * this.height] = v;
  }


  /*
   * Lua Script Interface
   */

  luaDebug(L) {
    const fb = lua.lua_touserdata(L, 1);
    console.log(fb);

    return 0;
  }

  luaSetpixel(L) {
    const fb = lua.lua_touserdata(L, 1);

    const x = lauxlib.luaL_checknumber(L, 2) - 1;
    const y = lauxlib.luaL_checknumber(L, 3) - 1;

    const r = lauxlib.luaL_optnumber(L, 4, 0.0);
    const g = lauxlib.luaL_optnumber(L, 5, 0.0);
    const b = lauxlib.luaL_optnumber(L, 6, 0.0);
    const w = lauxlib.luaL_optnumber(L, 7, 0.0);

    fb.setPixel(x, y, [r, g, b, w]);

    return 0;
  }

  luaGetpixel(L) {
    const fb = lua.lua_touserdata(L, 1);

    const x = lauxlib.luaL_checknumber(L, 2) - 1;
    const y = lauxlib.luaL_checknumber(L, 3) - 1;

    let pixel =  fb.getPixel(x, y);

    if (typeof(pixel) == 'undefined') {
      pixel = [0.0, 0.0, 0.0, 0.0];
    }

    lua.lua_pushnumber(L, pixel[0]);
    lua.lua_pushnumber(L, pixel[1]);
    lua.lua_pushnumber(L, pixel[2]);
    lua.lua_pushnumber(L, pixel[3]);

    return 4;
  }

  luaGetsize(L) {
    const fb = lua.lua_touserdata(L, 1);

    lua.lua_pushnumber(L, fb.width);
    lua.lua_pushnumber(L, fb.height);

    return 2;
  }
}


