pragma solidity ^0.5.0; // 不低于0.5 不高于 0.6

// 接口 , 其他合约可以调用 更像 strcut 一些 只是做声明
interface Regulator {
    //external 外部函数可以外部调用 regulator.checkValue() 这样子
    function checkValue(uint amount) external returns (bool);
    function loan() external returns (bool);
}
// 实现
contract Bank is Regulator {
    //private 仅当前合约能访问
    uint private value;

    constructor(uint amount) public {
        value = amount;
    }

    function deposit(uint amount) public {
        value += amount;
    }

    function withdraw(uint amount) public {
        if (checkValue(amount)) {
            value -= amount;
        }
    }

    function balance() public view returns (uint) {
        return value;
    }

    function checkValue(uint amount) public returns (bool) {
        // Classic mistake in the tutorial value should be above the amount
        return value >= amount;
    }

    function loan() public returns (bool) {
        return value > 0;
    }
}
// 实现
contract MyFirstContract is Bank(10) {
    string private name;
    uint private age;

    function setName(string memory newName) public {
        name = newName;
    }

    function getName() public view returns (string memory) {
        return name;
    }

    function setAge(uint newAge) public {
        age = newAge;
    }

    function getAge() public view returns (uint) {
        return age;
    }
}
