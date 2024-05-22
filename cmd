solc --optimize --abi --bin .\contract\Employee.sol -o build
abigen --abi=./build/Employee.abi --bin=./build/Employee.bin --pkg=api --out=./api/EmpRecord.go
solcjs --optimize --abi --bin .\Employee.sol -o build