'use client';

import { createPagesBrowserClient } from '@supabase/auth-helpers-nextjs';
import { createContext, useContext, useEffect, useState } from 'react';

const SupabaseContext = createContext<any>(null);

export const SupabaseProvider = ({ children }: { children: React.ReactNode }) => {
  const [supabase] = useState(() => createPagesBrowserClient());

  return (
    <SupabaseContext.Provider value={supabase}>
      {children}
    </SupabaseContext.Provider>
  );
};

export const useSupabase = () => useContext(SupabaseContext);
