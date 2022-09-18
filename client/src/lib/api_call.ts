export async function api_call(path: string) {
	let res = await fetch('http://localhost:80' + path)
	return await res.text()
}
