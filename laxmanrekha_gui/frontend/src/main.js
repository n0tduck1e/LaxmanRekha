import 'core-js/stable';
const runtime = require('@wailsapp/runtime');

// Main entry point
function start() {

	// Ensure the default app div is 100% wide/high
	var app = document.getElementById('app');
	app.style.width = '100%';
	app.style.height = '100%';

	// Inject html
	app.innerHTML = `
	<div class="box1">
	<div>
		<h1 class="heading">Laxman Rekha</h1>
	</div>
	<div class="input">
		<center>
		<table>
		<tr>
		<td><label>IP Address</label></td>
		<td><input id=ip ></input></td>
		</tr>
		<tr>
		<td><label>Username</label></td>
		<td><input id=user ></input></td>
		</tr>
		<tr>
		<td><label>Password</label></td>
		<td><input type=password id=pass ></input></td>
		</tr>
		<tr>
		<td><label>Port</label></td>
		<td><input id=port ></input></td>
		</tr>
		</table>
		<br>
		<br>
		<button class="glow-on-hover" id=submit>Start Configuration</button>
		</center>
	</div>
	</div>
	`;
	document.getElementById("submit").onclick = function() {
		var ip = document.getElementById("ip").value;
		var user = document.getElementById("user").value;
		var pass = document.getElementById("pass").value;
		var port = document.getElementById("port").value;
		window.backend.setup(ip,user,pass,port).then(result => { 
				app.innerHTML=`<div class="box1"><h2>All Done :)</h2></div>`;
		});

		app.innerHTML = ` <head>
			<style>
			.loading {
			position: absolute;
			top: 50%;
			left: 50%;
			margin: -15px 0 0 -15px;
			height: 30px;
			width: 30px;
			border: 2px solid #ddd;
			border-left-color: #009688;
			border-radius: 30px; /* border-radius: 50% */
			-webkit-animation: animation-rotate 950ms cubic-bezier(.64,2,.56,.6) infinite;
			animation: animation-rotate 950ms cubic-bezier(.64,2,.56,.6) infinite;
			}

			@-webkit-keyframes animation-rotate {
			100% {
			  -webkit-transform: rotate(360deg);
			}
			}
			@keyframes animation-rotate {
			100% {
			  -webkit-transform: rotate(360deg);
				  transform: rotate(360deg);
			}
			}
			</style>
			</head>
			<body>
			<div class="box1">
			<div class="loading"></div>
			<h2>Securing The Server</h2>
			</div>
			</body>
		`;

	};
};

// We provide our entrypoint as a callback for runtime.Init
runtime.Init(start);
