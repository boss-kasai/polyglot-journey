'use client';

import Link from 'next/link';

export default function HomePage() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen space-y-6">
      <h1 className="text-3xl font-bold">🎮 ゲームセンター</h1>
      <p className="text-lg">
        <ruby>
          遊<rt>あそ</rt>
        </ruby>
        びたいゲームを
        <ruby>
          選<rt>えら</rt>
        </ruby>
        んでね！
      </p>

      <div className="flex flex-col space-y-4">
        <Link href="/games/number-order">
          <button className="border border-blue-600 text-blue-600 px-4 py-2 rounded hover:bg-blue-50">
            🔢 <ruby>数字<rt>すうじ</rt></ruby>の<ruby>順番<rt>じゅんばん</rt></ruby><ruby>ゲーム<rt>げーむ</rt></ruby>
          </button>
        </Link>
      </div>
    </main>
  );
}
