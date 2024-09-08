// verifier.rs
use risc0_zkvm::Verifier;

pub fn verify_proof(proof: &[u8], public_input: &[u8]) -> Result<bool, Error> {
    let receipt = Verifier::deserialize(proof)?;
    Ok(receipt.verify(identity_circuit::ID, public_input))
}