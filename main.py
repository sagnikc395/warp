import click

@click.command()
@click.option("--count",default=1,help="Number of retries")
@click.option("--command",prompt="What is the command ?",help="The command to run")
def exec_command(count,command):
    '''
       simple program that execute the given command , given number of times  
    '''
    for _ in range(count):
        click.echo(f"hello from the otherside")





if __name__ == "__main__":
    exec_command()
