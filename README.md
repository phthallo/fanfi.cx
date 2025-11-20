# fanfi.cx
Query and read fanfiction from the Archive of Our Own over DNS. Inspired by [ch.at](https://github.com/Deep-ai-inc/ch.at). Runs on a single Go binary. 


## Usage
```
dig @fanfi.cx "your query here" TXT 
```

`+short` is optional but will make output look neater.

**Queries can and probably will time out, particularly with long works.** Add a timeout with `+timeout=<timeoutvaluehere>`. You might also need to rerun the command. I'm still not entirely sure why this happens. 

### Query terms

`[work_id] 22222` will fetch the work with the ID 22222. If used without `[chapter]`, it will fetch only the first chapter by default.

`[chapter] 3` can be used in conjunction with `[work_id]` and can be used for chapter-ination (pagination?? but for chapters??)

`[search] search query here` will search for that term. 

If no parameters are specified, it will default to searching for that term.

For instance, `dig @fanfi.cx "[work_id] 17400464 [chapter] 3" +short TXT` is a valid query.

`dig @fanfi.cx "[search] stag beetles and broken legs" +short TXT` is also a valid query, as is `dig @fanfi.cx "stag beetles and broken legs" +short TXT` 

## Development

1. Clone the repository
    ```
    git clone https://github.com/phthallo/fanfi.cx && cd fanfi.cx
    ```

2. Configure environment variables. For local development, you can use the following:
   ```
   FQDN="."
   PRIMARY_NS="ns1.hostmaster.com."
   SECONDARY_NS="ns2.hostmaster.com."
   TERTIARY_NS="ns3.hostmaster.com."
   QUATERNARY_NS="ns4.hostmaster.com."
   ```

3. Start the program.
    ```
    go run main.go
    ```

4. Test it!
    ```
    dig @0.0.0.0 "[search] your query" TXT +short 
    ```

## Production

Use the provided `docker-compose.yml` file in production. 

By default, the port used is port 53 - feel free to update this by adding `PORT=<yourport>` to your `.env`, though if you do this all `dig` queries will need to have `-p <yourport>` added on the end. 

Make sure you run `sudo ufw allow <yourport>` to open the port you use.


## Roadmap

- [ ] Search result pagination
- [ ] Tag/other metadata support for works in search view
- [ ] Overall work view from chapter 