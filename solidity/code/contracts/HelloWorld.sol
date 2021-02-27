// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract HelloWorld {
    string hello = "Hello World";
    function  getHello()  public view returns(string  memory){
        return hello;
    }
}
