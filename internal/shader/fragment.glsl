#version 440 core

in vec3 Color;
in vec2 TexCoord;

uniform sampler2D texture1;

out vec4 FragColor;

void main(){
    vec4 texColor=texture(texture1,TexCoord);
    FragColor=texColor*vec4(Color,1.)*2;
}