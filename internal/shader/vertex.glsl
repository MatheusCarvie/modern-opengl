// Versao do GLSL
#version 450 core

// Entrada / Saida
// layout(location=0)in vec3 aPos;
// layout(location=1)in vec3 aColor;
// layout(location=2)in vec2 aTexCoord;

// out vec3 ourColor;
// out vec2 TexCoord;
layout(location=0)in vec2 aPos;
layout(location=1)in vec3 aColor;

out vec3 ourColor;

void main(){
    ourColor=aColor;
    gl_Position=vec4(aPos,0.,1.);
}