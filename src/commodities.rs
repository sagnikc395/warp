pub struct Payment {
    amount: f64,
    sender: String,
    receiver: String,
}

impl Payment {
    fn new(amount:f64,sender:String,receiver:String) -> Self {
        Self {
            amount: amount,
            sender: sender,
            receiver: receiver,
        }
    }
}