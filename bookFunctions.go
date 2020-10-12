package main

func find(x string) int {
	for i, entry := range entries {
		if x == entry.IdTimeEntry {
			return i
		}
	}
	return -1
}

func updateEntry(id string, updatingEntry Entry){
	deleteEntry(id)
	entries = append(entries, updatingEntry)
}

func deleteEntry(id string){
	i := find(id)
	entries = append(entries[:i], entries[i+1:]...)
}


