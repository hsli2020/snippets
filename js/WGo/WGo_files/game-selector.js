(function(root) {

	'use strict';

	var Controller = {
		bindEvents : function bindEvents() {
			var gameModeEl = document.getElementById('gameMode');
			var root = domain + '.html';

			gameModeEl.addEventListener('change', function(ev) {
				var url = root + '?' + new Date().getTime() + '#' + ev.target.value;
				document.location.href = url;
			});
		},

		populateGameList : function populateGameList() {
			var container = document.getElementById('choose-game');
			var gameModeEl = document.getElementById('gameMode');
			var matchUrl = document.location.hash.substring(1);
			var html = (matchUrl.trim() === '') ? '<option value="">Click Here</option>' : '';

			matches.forEach(function (match) {
				html += '<option value="' + match.url + '">' + match.title + '</option>';
			});
			gameModeEl.innerHTML = html;
			container.classList.add('show');
		},

		setMatch : function setMatch() {
			var gameContainer = document.querySelector('.player');
			var matchUrl = document.location.hash.substring(1);
			var games = document.querySelectorAll('#gameMode option');

			if (matchUrl) {
				gameContainer.setAttribute('data-wgo', matchUrl);
				games = [].slice.call(games);

				games.forEach(function (game) {
					if (game.value === matchUrl) {
						game.setAttribute('selected', true);
					}
				});
				gameContainer.classList.add('show');
				return;
			}
		},

		init : function init() {
			Controller.populateGameList();
			Controller.setMatch();
			Controller.bindEvents();
		}
	};

	Controller.init();

}(window));
