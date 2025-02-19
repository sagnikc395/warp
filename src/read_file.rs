use std::{error::Error, io::Read};

use crate::consts::BUFFER_SIZE;

// snippet copied from here
// https://stackoverflow.com/questions/37079342/what-is-the-most-efficient-way-to-read-a-large-file-in-chunks-without-loading-th
pub fn read_file<R: Read>(mut reader: R) -> Result<(), Box<dyn Error>> {
    let mut buffer = [0_u8; BUFFER_SIZE];

    loop {
        let count = reader.read(&mut buffer)?;
        if count == 0 {
            break;
        }

        let string_slice = std::str::from_utf8(&buffer[..count])?;

        print!("{string_slice}");
    }

    Ok(())
}
