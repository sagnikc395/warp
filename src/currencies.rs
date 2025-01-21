// list of currencies that are valid that we need to check against in our
// .fin file and to typecheck and to return any other error

pub enum Currency {
    INR,
    USD,
    GBP,
    YUAN,
    //To add later
}


fn current_market_rate(from_currecy: Currency,to_currency:Currency)  {

}