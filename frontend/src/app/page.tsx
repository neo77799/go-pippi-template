'use client'

import { useState } from 'react'

export default function OmikujiPage() {
  const [result, setResult] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  const drawOmikuji = async () => {
    setLoading(true)
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/omikuji`)
      const data = await res.json()
      setResult(data.result)
    } catch (error) {
      console.error('おみくじ取得に失敗しました:', error)
      setResult('エラーが発生しました')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-4">
      <h1 className="text-3xl font-bold mb-8">おみくじアプリ</h1>
      
      <button
        onClick={drawOmikuji}
        disabled={loading}
        className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-400"
      >
        {loading ? '抽選中...' : 'おみくじを引く'}
      </button>

      {result && (
        <div className="mt-8 p-6 border-2 border-yellow-400 rounded-lg text-center">
          <h2 className="text-2xl font-bold mb-2">結果</h2>
          <p className="text-4xl font-bold text-red-500">{result}</p>
        </div>
      )}
    </div>
  )
}
