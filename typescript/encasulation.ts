//Bank Account
//Depositing
//Withdrawing
//Balance-hidden-encapsulated


class BankAccount{
    private _balance:number;
    constructor(initialbalance:number){
this._balance=initialbalance
    }
    //method deposite money
    public deposite(amount:number):void{
        if(amount<0){
            console.log("Invalid deposite amount")
            return

        }
        this._balance+=amount


    }
    public withdraw(amount:number):void{
        if(amount<0){
              console.log("Invalid deposite amount")
            return
        }
        if((this._balance-amount)<0){
            console.log("Insufficient Funds")
            return;
        }
        this._balance-=amount;


    }
    //Getter to get balance of the bank account
    public get balance():number{
        return this._balance
    }
}
const myAccount=new BankAccount(1000)

console.log(myAccount.balance)
myAccount.deposite(500)
myAccount.deposite(200)
console.log("Current Balance :", myAccount.balance)
