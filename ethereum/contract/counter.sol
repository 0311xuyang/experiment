pragma solidity ^0.4.0;
contract Counter {
    int private total = 0;
    mapping (string => int) private _bucket;

    function incrementCounter(string key) public {
        _bucket[key] += 1;
        total++;
    }
    
    function getCount(string key) public constant returns (int) {
        return _bucket[key];
    }
    
    function getTotal() public constant returns (int) {
        return total;
    }
}