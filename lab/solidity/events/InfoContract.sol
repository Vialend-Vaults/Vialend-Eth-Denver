pragma solidity ^0.8.10;
/*
abigen --abi=./InfoContract.abi --bin=./InfoContract.bin --pkg=api --out=./InfoContract.go
*/
contract InfoContract {
    
   string fName;
   uint age;
   
 	event Log(
	   address indexed,
	   uint256  indexed age,
       string name
    );
    
   function setInfo(string memory _fName, uint _age) public {
       fName = _fName;
       age = _age;
     
        emit Log(msg.sender, _age, _fName );
   }
   
   function getInfo() public view returns (string memory, uint) {
       return (fName, age);
   }   
}