# fanfi.cx
Query and read fanfiction from the Archive of Our Own over DNS. Inspired by [ch.at](https://github.com/Deep-ai-inc/ch.at). Runs on a single Go binary.


## Usage
```
dig @fanfi.cx "your query here" -p 1337 +short
```

`+short` is optional but will make output look neater.

### Query terms

`[work_id] 22222` will fetch the work with the ID 22222. If used without `[chapter]`, it will fetch only the first chapter by default.
`[chapter] 3` can be used in conjunction with `[work_id]` and can be used for chapter-ination (pagination?? but for chapters??)
`[search] search query here` will search for that term. 

If no parameters are specified, it will default to searching. for that term.

For instance, `dig @fanfi.cx "[work_id] 17400464 [chapter] 3" -p 1337 +short TXT` is a valid query. 
`dig @fanfi.cx "[search] stag beetles and broken legs" -p 1337 +short TXT` is also a valid query. 
