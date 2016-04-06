<template>
	<div class="row">
		<form class="col s12">
			<div class="row">
				<div class="input-field col s12">
					<textarea id="paste-area" class="materialize-textarea" v-model="pastedata"></textarea>
					<label for="textarea1">Paste your code</label>
				</div>
			</div>
		</form>
	</div>
	<div class="row">
		<div class="col s3 push-s9">
			<a class="waves-effect waves-light btn push-s7" v-on:click="paste">Paste!</a>
		</div>
	</div>
</template>

<script>
	module.exports = {
		data: function() {
			return {
				msg: 'Hello Tonnarello!',
				pastedata: ''
			}
		},
		methods: {
			paste: function() {
				var d = this.$get('pastedata');

				this.$http.post('/api/new', {data: d}).then(function(res) {
					console.log('Success!');
				}, function(err) {
					if (err) {
						console.error('Noes :-(');
					}
				});
			}
		}
	}
</script>
