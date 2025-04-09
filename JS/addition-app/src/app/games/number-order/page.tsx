'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useUser } from '@supabase/auth-helpers-react';
import LevelSelect from './LevelSelect';
import StartScreen from './StartScreen';
import PlayScreen from './PlayScreen';
import ResultScreen from './ResultScreen';
import { PositionedNumber, LevelOption } from '../../../lib/gameUtils';

export default function GamePage() {
  const user = useUser();
  const router = useRouter();
  const [checkingAuth, setCheckingAuth] = useState(true);

  const [level, setLevel] = useState<LevelOption | null>(null);
  const [numbers, setNumbers] = useState<PositionedNumber[]>([]);
  const [started, setStarted] = useState(false);
  const [startTime, setStartTime] = useState<number | null>(null);
  const [endTime, setEndTime] = useState<number | null>(null);

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

  const resetToLevelSelect = () => {
    setLevel(null);
    setStarted(false);
    setStartTime(null);
    setEndTime(null);
  };

  if (!level) {
    return (
      <div className="relative w-full h-[80vh]">
        <h1 className="text-2xl font-bold text-center mt-6">
          🔢 <ruby>数字<rt>すうじ</rt></ruby>タッチゲーム
        </h1>
        <LevelSelect onSelectLevel={setLevel} />
      </div>
    );
  }

  if (!started) {
    return (
      <div className="relative w-full h-[80vh]">
        <h1 className="text-2xl font-bold text-center mt-6">
          🔢 <ruby>数字<rt>すうじ</rt></ruby>タッチゲーム
        </h1>
        <StartScreen
          level={level}
          onStart={(nums) => {
            setNumbers(nums);
            setStarted(true);
            setStartTime(Date.now());
            setEndTime(null);
          }}
        />
      </div>
    );
  }

  if (endTime && startTime) {
    return (
      <ResultScreen
        startTime={startTime}
        endTime={endTime}
        onRetry={() => {
          setStarted(false);
          setEndTime(null);
        }}
        onBack={resetToLevelSelect}
      />
    );
  }

  return (
    <div className="relative w-full h-[80vh]">
      <h1 className="text-2xl font-bold text-center mt-6">
        🔢 <ruby>数字<rt>すうじ</rt></ruby>タッチゲーム
      </h1>
      <PlayScreen
        numbers={numbers}
        startTime={startTime}
        onClear={(end) => setEndTime(end)}
      />
    </div>
  );
}
