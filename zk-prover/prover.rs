// prover.rs
use risc0_zkvm::{Prover, ProverOpts};

pub fn generate_proof(identity: &Identity) -> Result<Vec<u8>, Error> {
    let mut prover = Prover::new_with_opts(ProverOpts::default())?;
    prover.add_input_u8_slice(identity.as_bytes());
    let receipt = prover.run()?;
    Ok(receipt.serialize())
}