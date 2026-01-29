const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'

export async function getFundamentals(ticker: string) {
	// TODO: call internal API route and return parsed payload
	try {
		const res = await fetch(`${API_BASE}/fundamentals/${ticker}`)
		if (!res.ok) throw new Error('Failed to fetch fundamentals')
		return await res.json()
	} catch (error) {
		console.error('getFundamentals error:', error)
		throw error
	}
}

export async function listTickers() {
	// TODO: call backend list endpoint
	try {
		const res = await fetch(`${API_BASE}/tickers`)
		if (!res.ok) throw new Error('Failed to fetch tickers')
		return await res.json()
	} catch (error) {
		console.error('listTickers error:', error)
		throw error
	}
}

export async function getSwaggerSpec() {
	// TODO: fetch swagger.json
	try {
		const res = await fetch(`${API_BASE}/../swagger.json`)
		if (!res.ok) throw new Error('Failed to fetch swagger spec')
		return await res.json()
	} catch (error) {
		console.error('getSwaggerSpec error:', error)
		throw error
	}
}
