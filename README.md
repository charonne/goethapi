<h1>GoethApi</h1>

<h1>Installation</h1>

<h2>Config file</h2>
In order to set your own config, you have to create your config file.<br />
rename <code>config/app-example.yml</code> to <code>config/app.yml</code><br /> and change the config to your own settings<br />



<h1>Todo</h1>





<h1>Examples</h1>

Url: http://localhost:8080/contract/create/
Method: Post
Data: {"name": "simplestorage", "source": "test code etc."}

Url: http://localhost:8080/contract/deploy/
Method: Post
Data: {"id": "<contract.idkey>", "params": ["42"]}

Url: http://localhost:8080/contract/exec/
Method: Post
Data: {"address": "<contract.idkey>", "method": "Set", "params": ["42"]}
{"address": "g73xKOwc", "method": "Get", "params": []}
