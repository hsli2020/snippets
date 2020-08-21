pragma solidity ^0.4.21;

contract TruthVote {
    enum States {
        REGISTER,
        VOTE,
        DISPERSE
    }

    address public owner = msg.sender;

    uint voteCost;

    address[] trueVotes;
    address[] falseVotes;

    mapping (address => bool) voters;
    mapping (address => bool) hasVoted;

    uint VOTE_COST = 100;

    States state;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyVoter() {
        require(voters[msg.sender] != false);
        _;
    }

    modifier isCurrentState(States _stage) {
        require(state == _stage);
        _;
    }

    modifier hasNotVoted() {
        require(hasVoted[msg.sender] == false);
        _;
    }

    function startVote()
        public
        onlyOwner()
        isCurrentState(States.REGISTER)
    {
        goToNextState();
    }

    function goToNextState() internal {
        state = States(uint(state) + 1);
    }

    modifier pretransition() {
        goToNextState();
        _;
    }

    function addVoter(address voter)
        public
        onlyOwner()
        isCurrentState(States.REGISTER)
    {
        voters[voter] = true;
    }

    function vote(bool val)
        public
        payable
        isCurrentState(States.VOTE)
        onlyVoter()
        hasNotVoted()
    {
        if (msg.value >= VOTE_COST) {
            if (val) {
                trueVotes.push(msg.sender);
            } else {
                falseVotes.push(msg.sender);
            }
            hasVoted[msg.sender] = true;
        }
    }

    function disperse(bool val)
        public
        onlyOwner()
        isCurrentState(States.VOTE)
        pretransition()
    {
        address[] memory winningGroup;
        uint winningCompensation;
        if (trueVotes.length > falseVotes.length) {
            winningGroup = trueVotes;
            winningCompensation = VOTE_COST + (VOTE_COST*falseVotes.length) / trueVotes.length;
        } else if (trueVotes.length < falseVotes.length) {
            winningGroup = falseVotes;
            winningCompensation = VOTE_COST + (VOTE_COST*trueVotes.length) / falseVotes.length;
        } else {
            winningGroup = trueVotes;
            winningCompensation = VOTE_COST;
            for (uint i = 0; i < falseVotes.length; i++) {
                falseVotes[i].transfer(winningCompensation);
            }
        }

        for (uint j = 0; j < winningGroup.length; j++) {
            winningGroup[j].transfer(winningCompensation);
        }
    }
}
