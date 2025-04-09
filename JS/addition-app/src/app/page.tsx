'use client';

import Link from 'next/link';
import { useEffect, useState } from 'react';
import { useUser } from '@supabase/auth-helpers-react';
import { useRouter } from 'next/navigation';

export default function HomePage() {
  const user = useUser();
  const router = useRouter();
  const [checkingAuth, setCheckingAuth] = useState(true);

  useEffect(() => {
    if (user === null) {
      router.replace('/login');
    } else {
      setCheckingAuth(false);
    }
  }, [user, router]);

  if (checkingAuth) {
    return <p className="text-center mt-20">ログイン状態を確認中...</p>;
  }

  return (
    <main className="flex flex-col items-center justify-center min-h-screen space-y-8 bg-blue-50 text-center px-4">
      <h1 className="text-4xl font-bold text-pink-600">
        🎮 <ruby>ゲーム<rt>げーむ</rt></ruby>センター
      </h1>

      <p className="text-xl text-gray-800">
        <ruby>遊<rt>あそ</rt></ruby>びたい
        <ruby>ゲーム<rt>げーむ</rt></ruby>を
        <ruby>選<rt>えら</rt></ruby>んでね！
      </p>

      <div className="flex flex-col space-y-6 w-full max-w-xs">
        <Link href="/games/number-order">
          <button className="bg-yellow-300 text-gray-800 text-xl font-bold px-6 py-6 rounded-full shadow-md hover:bg-yellow-400 transition leading-normal">
            🔢 <ruby>数字<rt>すうじ</rt></ruby>の
            <ruby>順番<rt>じゅんばん</rt></ruby>
            <ruby>ゲーム<rt>げーむ</rt></ruby>
          </button>
        </Link>
      </div>
    </main>
  );
}
