import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
	// TODO: validate ticker param and call backend API
	res.status(200).json({ message: 'fetchFundamentalsAPI placeholder' })
}

export function validateTickerParam(req: NextApiRequest) {
	// TODO: implement validation
	return true
}

export function cacheFundamentals(ticker: string, payload: any) {
	// TODO: implement caching layer
	return
}