
import CodeMirror from 'codemirror'
import Runner from './engine/runner'
import Viewport from './engine/viewport'

import CodeMirrorLua from 'codemirror/mode/lua/lua'

let runner = new Runner();

let textarea = document.getElementById("shadercode");
let editor = CodeMirror.fromTextArea(textarea, {
  lineNumbers: true,
  mode: "lua",
  theme: "monokai",
});

runner.loadScript(editor.getValue());
runner.render();


let canvas = document.getElementById("shaderpreview");
let vp = new Viewport(canvas);

vp.render(runner.framebuffer);


