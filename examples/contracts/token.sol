pragma solidity ^0.4.8;

contract CoinFidToken {
    /* Public variables of the token */
    string public standard = 'Token 0.1';
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
    uint256 public totalDistributedCoins;
    uint256 public totalRewardsOffered;

    /* This creates array variables of the token */
    mapping (address => bool) public superAdmins;           // A superAdmin is a CoinFid employee
    mapping (address => bool) public admins;                // An admin is a company (e.g. Decathlon)
    mapping (address => bool) public distributors;          // A distributor is a device owner of the company (e.g. Decathlon employee)
    mapping (address => uint256) public balanceOf;          // This creates an array with all balances


    /* This generates a public event on the blockchain that will notify clients */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /* This notifies clients about the amount burnt */
    event Burn(address indexed from, uint256 value);

    /* Initializes contract with initial supply tokens to the creator of the contract */
    function CoinFidToken(
        uint256 initialSupply,
        string tokenName,
        uint8 decimalUnits,
        string tokenSymbol,
        address firstSuperAdmin
        ) {
        totalSupply = initialSupply;                        // Update total supply
        name = tokenName;                                   // Set the name for display purposes
        symbol = tokenSymbol;                               // Set the symbol for display purposes
        decimals = decimalUnits;                            // Amount of decimals for display purposes
        totalDistributedCoins = 0;                          //
        totalRewardsOffered = 0;                            //

        superAdmins[firstSuperAdmin] = true;                // Define the first superAdmin
        admins[firstSuperAdmin] = true;                     //
        distributors[firstSuperAdmin] = true;               //

        admins[msg.sender] = true;                          // Define the creator as admin
        distributors[msg.sender] = true;                    //
        balanceOf[msg.sender] = initialSupply;              // Give the creator all initial tokens e.g. 0 tokens
    }

    /* Update Coin parameters */
    function setName(string _newTokenName) {
        name = _newTokenName;                               // Update token name
    }

    function setTokenSymbol(string _newTokenSymbol) {
        symbol = _newTokenSymbol;                           // Update token symbol
    }


    /* Provide and remove rights */
    function nameSuperAdmin(address _newSuperAdmin) {
        require(superAdmins[msg.sender] == true);           // Check if the sender is a superAdmin
        superAdmins[_newSuperAdmin] = true;                 // Declare a new superAdmin address
        admins[_newSuperAdmin] = true;                      //
        distributors[_newSuperAdmin] = true;                //
    }

    function nameAdmin(address _newAdmin) {
        require(admins[msg.sender] == true);                // Check if the sender is an admin
        admins[_newAdmin] = true;                           // Declare a new admin address
        distributors[_newAdmin] = true;                     //
    }

    function nameDistributor(address _newDistributor) {
        require(admins[msg.sender] == true);                // Check if the sender is an admin
        distributors[_newDistributor] = true;               // Declare a new distributor address
    }

    function removeSuperAdmin(address _superAdmin) {
        require(superAdmins[msg.sender] == true);           // Check if the sender is a superAdmin
        require(msg.sender != _superAdmin);                 // Check if the superAdmin is not removing his own rights
        superAdmins[_superAdmin] = false;                   // Remove the superAdmin rights
    }

    function removeAdmin(address _admin) {
        require(admins[msg.sender] == true);                // Check if the sender is an admin
        require(msg.sender != _admin);                      // Check if the admin is not removing his own rights
        admins[_admin] = false;                             // Remove the admin rights
        superAdmins[_admin] = false;                        // Make sure the former admin had no superAdmin rights
    }

    function removeDistributor(address _distributor) {
        require(admins[msg.sender] == true);                // Check if the sender is an admin
        require(msg.sender != _distributor);                // Check if the admin is not removing his own rights
        distributors[_distributor] = false;                 // Remove the distributor rights
        admins[_distributor] = false;                       // Make sure the former distributor had no admin rights
        superAdmins[_distributor] = false;                  // Make sure the former distributor had no superAdmin rights
    }



    /* Distribute coins to customers - Method for distributors only */
    function distribute(address _to, uint256 _value) {
        require(_to != 0x0);                                // Prevent transfer to 0x0 address. Use burn() instead
        require(distributors[msg.sender] == true);          // Check if the sender is a distributor
        require(_to != msg.sender);                         // Verifiy the sender is not sending coins to himself
        balanceOf[msg.sender] += _value;                    // Load sender balance
        totalSupply += _value;                              // Updates totalSupply
        totalDistributedCoins += _value;                    // Increment counter of distributed coins

        transfer(_to, _value);                              // Send coins to recipient
    }

    /* Send coins */
    function transfer(address _to, uint256 _value) {
        require(_to != 0x0);                                // Prevent transfer to 0x0 address. Use burn() instead
        require(balanceOf[msg.sender] >= _value);           // Check if the sender has enough
        require(balanceOf[_to] + _value >= balanceOf[_to]); // Check for overflows
        balanceOf[msg.sender] -= _value;                    // Subtract from the sender
        balanceOf[_to] += _value;                           // Add the same to the recipient
        Transfer(msg.sender, _to, _value);                  // Notify anyone listening that this transfer took place
    }

    /* Distribute coins to customers - Method for distributors only */
    function generate(uint256 _value) {
        require(superAdmins[msg.sender] == true);           // Check if the sender is a superAdmin
        balanceOf[msg.sender] += _value;                    // Load sender balance
        totalSupply += _value;                              // Updates totalSupply
    }

    function changeCoinsForReward(address _from, uint256 _value){
        require(distributors[msg.sender] == true);          // Check if the sender is a distributor
        require(balanceOf[_from] >= _value);                // Check if the sender has enough
        balanceOf[_from] -= _value;                         // Subtract from the sender
        totalSupply -= _value;                              // Updates totalSupply
        if(distributors[_from] != true){                    // Increment counter of rewards offered only if the recipient is a customer
            totalRewardsOffered +=1;
        }
        Burn(_from, _value);
    }

    function balanceOf(address _address) constant returns (uint256){
        return(balanceOf[_address]);
    }

}
