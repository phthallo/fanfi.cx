# fanfi.cx

Query and read fanfiction from the Archive of Our Own over DNS. Inspired by [ch.at](https://github.com/Deep-ai-inc/ch.at). Runs on a single Go binary. 


![demo](https://github.com/user-attachments/assets/39cf586a-441a-41de-a24f-7b49a99806ee)
<p align = "center"><i>if there's a <s>screen</s> terminal ao3 shall be seen</i></p>



## Usage
```
dig @fanfi.cx "test" TXT +tcp +short
```

Run the above command in your terminal. `+short` is optional but will make output look neater. `+tcp` is also optional, but I recommend it because queries usually time out without it.

### Query terms

`[work_id] 17400464` will fetch the work with the ID 17400464. If used without `[chapter]`, it will fetch only the first chapter by default.

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

You can use the provided `docker-compose.yml` file in production, or just build it yourself.

By default, the port used is port 53 - feel free to update this by adding `PORT=<yourport>` to your `.env`, though if you do this all `dig` queries will need to have `-p <yourport>` added on the end. 

Make sure you run `sudo ufw allow <yourport>` to open the port you use.

## Roadmap

- [ ] Search result pagination
- [ ] Tag/other metadata support for works in search view
- [ ] Overall work view from chapter
- [ ] Rewrite of the parameter interpretation function.

## Notes
- If the work has been deleted, can only be viewed by logged-in users only, or is otherwise restricted, you won't be able to access it using this tool.
- I do not hold any ownership or responsibility over the content you might see when you use this tool. Seriously. Here be dragons etc etc. 

## why 💀
1. it's silly
2. you can read on plane wifi now 
3. yes 
