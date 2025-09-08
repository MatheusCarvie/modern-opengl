#version 450 core

in vec3 ourColor;

out vec4 aColor;

void main(){
    aColor=vec4(ourColor,1.);
}