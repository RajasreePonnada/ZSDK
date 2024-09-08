// avail_da_client.rs
use avail_sdk::{AvailClient, SubmitDataRequest};

pub async fn submit_identity_data(client: &AvailClient, data: Vec<u8>) -> Result<Hash, Error> {
    let request = SubmitDataRequest {
        data,
        app_id: Some(1), // Assuming 1 is the app ID for identity management
    };

    let hash = client.submit_data(request).await?;
    Ok(hash)
}