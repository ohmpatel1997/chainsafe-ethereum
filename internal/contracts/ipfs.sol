pragma solidity >=0.7.0 <=0.8.3;

contract IPFS {

    string[] cIds;


    function printIPFSIdentifiers() // defining a function
    public // this function is callable by anyone
    view // dictates that this function promises to not modify the state
    returns (string[] memory) // function returns a string variable from memory
    {
        return cIds;
    }

    function setIPFSIdentifier(string memory _cid)
    public

    {
        uint arrayLength = cIds.length;
        for (uint i=0; i<arrayLength; i++) {

            if (keccak256(bytes(cIds[i])) == keccak256(bytes(_cid))) {
                return;
            }
        }
        cIds.push(_cid);
    }
}