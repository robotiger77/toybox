components {
  id: "hook"
  component: "/main/hook.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"hook\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
  position {
    y: 11.0
    z: 0.1
  }
}
