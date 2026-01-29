import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
	if (req.method !== 'GET') {
		res.status(405).json({ error: 'Method not allowed' })
		return
	}

	// TODO: validate ticker param and call backend API
	res.status(200).json({ message: 'fetchFundamentalsAPI placeholder' })
}

export function validateTickerParam(req: NextApiRequest): boolean {
	// TODO: implement validation
	return true
}

export function cacheFundamentals(ticker: string, payload: any): void {
	// TODO: implement caching layer
	return
}