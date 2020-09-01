// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
        fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
        for i, s := range slice {
                fmt.Printf("[%d] %p %s\n",
                        i,
                        &slice[i],
                        s)
        }
}
