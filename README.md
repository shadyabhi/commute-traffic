# commute-traffic

It's a WIP.

## Config file

	apiKeys:
		# Get API key from https://developers.google.com/maps/documentation/directions/get-api-key
		- secret-api-key

	tracks:
		- source: Zonasha Elegance, Harlur Main Road
		  destination: LinkedIn, Global Technology Park
		- source: Zonasha Elegance, Harlur Main Road
		  destination: Slurp Cafe, Indiranagar

	elasticsearch:
		indexName: commute-traffic
		docType: doc

## Run

Once it's installed, you can run this as a periodic job. A sample run looks like:-

	➜ $?=0 @arastogi-ld2 repos/commute-traffic [12:45AM] (master|…)➤ ./commute-traffic
	2017/11/14 00:46:06 Currently working on, SRC=Zonasha Elegance, DST=LinkedIn, Global Technology Park
	2017/11/14 00:46:07 Currently working on, SRC=Zonasha Elegance, DST=Slurp Cafe, Indiranagar
	2017/11/14 00:46:07 Indexed to ES: doc={%!s(int64=6787) 1510649166 %!s(int64=1085) Zonasha Elegance, Reliable Silver Oak Layout, Eastwood Township, Harlur, Bengaluru, Karnataka 560035, India LinkedIn, Global Technology Park, 7th floor,Tower A, Global Technology Park, Devarabisanahalli, Next to Intel, Adarsh Palm Retreat, Bengaluru, Karnataka 560103, India %!s(int64=1168)}
	2017/11/14 00:46:07 Indexed to ES: doc={%!s(int64=14380) 1510649167 %!s(int64=2477) Zonasha Elegance, Reliable Silver Oak Layout, Eastwood Township, Harlur, Bengaluru, Karnataka 560035, India Slurp Cafe, Indiranagar, 1079, 12th Main Road, HAL 2nd Stage, Indiranagar, Bengaluru, Karnataka 560038, India %!s(int64=2603)}

	>>>  1s elasped...
	➜ $?=0 @arastogi-ld2 repos/commute-traffic [12:46AM] (master|…)➤
