# HODL.ZONE

## Abstract

https://hodl.zone is a site that provides cryptocurrency owners a means of earning interest on their holdings. The name itself refers to a popular intentional mispelling on "hold", refering to not selling. Hodl.zone uses Poloniex's margin trading system to earn interest, by providing loans to traders. These loans are backed by Poloniex, and therefore have no risk of defaulting. 

Hodl.zone manages a lending bot that runs on an algorithm built in house to maxmimize lending uptime, and lending rates. Our users are provided with a suite of graphs and information to track their earnings. We provide the suite to not only provide customers with daily personal updates, but also maintain the payment side of the platform. Hodl.zone for part of it's operation was charging up to 10% on profits made by our users. A credit system was credited, as well as a referral program and other means of reducing the fees we charged.


# Why I created it

https://Hodl.zone is one of my most ambitious projects, and not only by lines of code. I've been interested in cryptocurrency for years, and at one point I became aware of the Poloniex Lending system. This enabled customers on the site provide loans (which were often as short as a few hours) to margin traders. This realization spurred a lot of research, and subsequently experimentation. I created a rough lending bot that was naive, and made very poor lending rate choices. There was no risk to setting a loan at 0.001% per day, but there wasn't much of a benefit. 

I wrote another another script to track the lending rates, and realized my naive lending algorithm was vastly under performing. I made a better algorithm, and was making enough money in interest to cover lunch money. When a buddy of mine got some cryptocurrency, he was interested in what I made. I was runnning 3 instances of the bot for friends at that point, and decided to write a system to accomadate all of us. My friend joined me, and what started as a weekend project, turned into months of work.

Hodl.zone was created, however we prevented random signups by an access code system. We got immediate customers who wanted in, and for the first 3 months of operation, we averaged 7% interest per month. Unfortunately the lending market did not hold like that forever, and 0.8-1.5% monthy was more common. 

We finally shut our doors when Poloniex removed margin trading support for people in the US. Most of our customers were in Asia (hence why the site has a chinese and tawainese translation), however once we made it free for our customers, it no longer paid for itself.

# Project Components

## Website Frontend

The frontend is created mostly with custom html/css/js. We used a framwork for the user dashboard, which cannot be reached anymore.

The frontend supports multiple languages (US, Chinese, and Tawainese) to support our diverse community. The user dashboard contained many graphs to reflect profits and lending rates. We also provided users with a log system that exports logs related to their account to a page they can access. We provided this when users would ask questions about issues they were experiencing, mostly due to mis configuration. We provided users with configuration settings that would affect the lending bot when used on their account, and many of the errors could be easily corrected when providing them with the logs.

## Database

The Database used grew as we grew our platform. We started with an embeded key/val leveldb, but switched to mongodb for most of our data. We decided to go with a document store over a more structured relational db (sql), as we were unsure of what fields we were to include. Mongodb more easily supported our very agile development method, and made database administration almost non existant, as there were no database schemas to mess around with.

### User Database

Our user database was split into subsections. The first obviously being the authentication. All passwords were stored as a salted hash, and some extra data was also stored here like 2 factor authentication seeds, login settings (auto log after X minutes), etc.

The user payment section kept track of user deposits and charges. We used coinbase for payment recieving, and had to recieve callbacks from coinbase to notify us of payment. Correct tracking of payments was critical, and was created to be indepotent to allow for manual intervention, without messing up the automatic system.

The last section was the user data section. We tracked ours users earnings in great detail, to provide information, and for our auditing purposes. Each time a loan closed, we had to calculate the charge. To ensure our users had all the information to dispute a charge (and us all the information to justify one), we had all data easily to navigate. Because of our limited disk space, a lot of the extra data was removed after 1 month.

### Poloniex Lending Database

The lending bot algorithm was improved to use historical data in its decision making. By using historial data, the bot could make optimize lending uptime vs lending rates. For this reason, we had a database that stored all lending rates published by Poloniex, and rates published by the bot itself. The bot would frquently "test the books", to measure volume, as volume was published. This would allow the bot to not post the lowest rate, but post some percentage higher than the rates on the books, and still get taken in a reasonable amount of time.

## Lending Bot

The lending bot was the heart of hodl.zone. If it failed, so did the site. It was upmost important that it not only worked, but it performed well. If users figured they could do better on their own, the site would also fail. A lot of testing and measuring was done to tune this, and rates were manually analyzed frequently to adjust any tuning. 

**The problems**. The bot did not come without problems, the first major one being rate limiting. As our userbase grew, the number of api requests we had to send also grew. Poloniex rate limits per IP pretty heavily. To prevent accounts from being locked, we had to use the same IP for an account, preventing the use of common proxy tools. We took advantage of very cheap, very lightweight cloud machines, and had our central machine push rate the lightweight boxes. This was our Master + Slave model (referred to as the hive and bees, as slaves felt a bit violent).

### Master Bot

The master bot served 2 purposes. Find the best lending rates at any given moment, and manage the user accounts on each slave. It wants to keep the amount of users on each slave balanced, and would handle rebalancing in the case of a slave crashing/coming online. This autoscaling feature was to make the system robust, as we don't want to play sys admins.

### Slave Bots

Slave bots recieve rates from the master, and then post loans for the accounts they control. The master can also tell a slave to drop/add a user. The slave manages it's rate limit, and reports frequent hearbeats about which users it is controlling, any errors, and the last api request timestamp for each account.

### Auditor Bot

When creating the master/slave system, it came apparent that is was pretty complex, and therefore a bug was almost guarenteed. My biggest fear, would to have an account fall between the cracks, and not recieve loans. To prevent this, a 3rd component was created. The Auditor. The auditor's job was to run an audit every hour. It would:
- Check to make sure all accounts were managed, and had an api request in the last hour.
- Each slave the master is using is actually online
- The balance of users was reasonable

Given the event a condition was not met, it had the power to correct it. If the errors persisted, it would notify us via an alert. Each audit left behind a document that served as an audit report. We could also trigger these audits manually. This was our primary eye into the lending system from a user management point of few, and worked like a charm.

## Monitoring

Aside from lending monitoring, we employed a grafana/prometheus monitoring system to track all of our cloud machines. We had a slack alerting system if any component was non-responsive.

## Client interactions

Hodl.zone was more than a technical project, it was also a social one. We had a client base, and had to keep them in the loop and take their feedback/questions/concerns/problems. We managed all conversations in a slack channel, this was the best way to get the quickest support. We did create an email setup to provide us with professional emails (steven@hodl.zone, support@hodl.zone, etc), but was used less frequently than the slack.

For all major updates and announcements, if a slack post was not enough we used Medium. This stored our getting started guides, and major updates that required user interaction. (https://medium.com/hodlzone)
