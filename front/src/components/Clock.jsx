import { useState, useEffect } from "react";

function Clock() {
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const id = setInterval(() => {
      setTime(new Date());
    }, 1000);
    return () => clearInterval(id);
  }, []);

  return (
    <div className="flex flex-col justify-center items-center gap-2">
      <span className="text-2xl	font-jamsil">
        {time.getFullYear()}-{String(time.getMonth() + 1).padStart(2, 0)}-
        {time.getDate()}
      </span>
      <span className="text-2xl	font-jamsil">{time.toLocaleTimeString()}</span>
    </div>
  );
}

export default Clock;
