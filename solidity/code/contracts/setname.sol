pragma solidity ^0.5.0;


contract SetNameContract{
    //private 仅在当前合约能被访问,继承的合约不可
    string private name;
    uint private age;
    //public 让函数在当前合约,和被继承的合约内都可以访问
    //memory 算是开辟一个内存
    function setName(string memory newName) public{
        name = newName;
    }
    //view 可读不可修改
    function getName() public view returns (string memory){
        return name;
    }
}
