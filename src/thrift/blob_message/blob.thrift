struct Message {
    1: string partition_id
    2: string blob_id
}

service BlobSvc {
    /// requests <partition_id, blob_id>
    /// returns a server message
    string request(1: Message msg )
}
