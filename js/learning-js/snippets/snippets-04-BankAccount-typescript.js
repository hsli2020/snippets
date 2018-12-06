// BankAccount.ts
class BankAccount {  
    constructor(public balance: number) {  
    }  
    deposit(credit: number) {  
        this.balance += credit;  
        return this.balance;  
    }  
}

class CheckingAccount extends BankAccount {  
    constructor(balance: number) {  
        super(balance);  
    }  
    writeCheck(debit: number) {  
        this.balance -= debit;  
    }  
}


// BankAccount.js
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};

var BankAccount = (function () {
    function BankAccount(balance) {
        this.balance = balance;
    }
    BankAccount.prototype.deposit = function (credit) {
        this.balance += credit;
        return this.balance;
    };
    return BankAccount;
}());

var CheckingAccount = (function (_super) {
    __extends(CheckingAccount, _super);
    function CheckingAccount(balance) {
        _super.call(this, balance);
    }
    CheckingAccount.prototype.writeCheck = function (debit) {
        this.balance -= debit;
    };
    return CheckingAccount;
}(BankAccount));
