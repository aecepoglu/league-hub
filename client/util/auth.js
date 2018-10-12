import store from "local-storage";

function save(token) {
	store.set("token", token);
}

function load() {
	return store.get("token");
}

function clear() {
	store.clear();
}

export default { clear, load, save };
