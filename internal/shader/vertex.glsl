// Versao do GLSL
#version 440 core

// Entrada / Saida
// layout(location=0)in vec3 aPos;
// layout(location=1)in vec3 aColor;
// layout(location=2)in vec2 aTexCoord;

// out vec3 ourColor;
// out vec2 TexCoord;
layout(location=0)in vec2 aPos;
layout(location=1)in vec3 aColor;
layout(location=2)in vec2 aTexCoord;

out vec2 TexCoord;
out vec3 Color;

void main(){
    Color=aColor;
    TexCoord=aTexCoord;
    gl_Position=vec4(aPos,0.,1.);
}