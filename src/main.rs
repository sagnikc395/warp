use clap::{command, Parser};

#[derive(Parser, Debug)]
#[command(author,version,about,long_about = None)]
struct Args {
    //openai/claude API key
    #[arg(short = 'k', long)]
    api_key: String,

    //local doc name
    #[arg(short, long)]
    limit: u16,

    #[arg(short, long)]
    name: String,
}

fn main() {
    let args = Args::parse();
    println!("Hello {} !", args.name);
}
