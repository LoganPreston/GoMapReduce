# Map Reduce Filter

This is a simple map reduce filter implementation with generics and simple testing. There are two implementations of filter because I was curious how much more performant array accesses would be than appends. They're actually closer than I expected, leading me to think appending is pretty efficient with adequately sized slices.
