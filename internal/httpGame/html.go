package httpGame

const newHtml string = `<html>
<head><title>Tic-tac-toe</title></head>
<body>
<div class="config">
	<form>
		<div>Config new game</div>
		<div>Setup player side:
			<input type="radio" name="side" value="X" checked/>
			<label for="side">X</label>
			<input type="radio" name="side" value="O"/>
			<label for="side">O</label>
		</div>
		<div>Setup field size:
			<input type="text" name="fieldSize" value="3" />
		</div>
		<div>Setup win sequence length:
			<input type="text" name="winSeq" value="3" />
		</div>
		<div><button type="send">Start new game</button></div>
	</form>
</div>
</body>
</html>`
