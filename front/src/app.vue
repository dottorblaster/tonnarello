<template>
	<div class="row">
		<form class="col s12">
			<div class="row">
				<div class="input-field col s12">
					<textarea id="paste-area" class="materialize-textarea"></textarea>
					<label for="textarea1">Paste your code</label>
				</div>
			</div>
		</form>
	</div>
	<div class="row">
		<div class="col s3 push-s9">
			<a class="waves-effect waves-light btn push-s7">Paste!</a>
		</div>
	</div>
</template>

<script>
	module.exports = {
		data: function() {
			return {
				msg: 'Hello Tonnarello!'
			}
		},
		methods: {
			paste: function() {}
		}
	}
</script>
