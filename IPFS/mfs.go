/* MFS */

// copy a single file into a directory
await ipfs.files.cp('/example-file.txt', '/destination-directory')
await ipfs.files.cp('/ipfs/QmWGeRAEgtsHW3ec7U4qW2CyVy7eA2mFRVbk1nb24jFyks', '/destination-directory')

// copy multiple files into a directory
await ipfs.files.cp('/example-file-1.txt', '/example-file-2.txt', '/destination-directory')
await ipfs.files.cp('/ipfs/QmWGeRAEgtsHW3ec7U4qW2CyVy7eA2mFRVbk1nb24jFyks',
 '/ipfs/QmWGeRAEgtsHW3jk7U4qW2CyVy7eA2mFRVbk1nb24jFyre', '/destination-directory')

// copy a directory into another directory
await ipfs.files.cp('/source-directory', '/destination-directory')
await ipfs.files.cp('/ipfs/QmWGeRAEgtsHW3ec7U4qW2CyVy7eA2mFRVbk1nb24jFyks', '/destination-directory')
