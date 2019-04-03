# Blockchain Based Coin Raffle

## Abstract

As a fun community event, a company that I am apart of called "The Factoid Authority" decided to host a free raffle. The prize was a limited edition silver coin etched with the Factom logo, Factom being a blockchain technology. In the spirit of the blockchain, the coin raffle was decided to be hosted on chain, and therefore 100% transparent and auditable.

## Design

In order to enter the raffle, a user would submit a post on a designed thread on https://factomize.com. This forum is unique, as all posts are also recorded on the blockchain (making the forum immutable). The post on the blockchain has an essentially random transaction id. To prevent users from "mining" a txid, I would make a closing post after 24 hrs. My entry's hash would be concatenated with theirs, preventing any ranking of the votes prior to my ending post. 

We would then simply award the prize to the lowest hash of the concatenated string. The community had a lot of fun with this, and its use of blockchain was perfect to ensure a transparent contest.

