import Clock from "../components/Clock";

const Main = () => {
  return (
    <div className="w-screen h-screen bg-zinc-200	flex flex-col	justify-center items-center gap-5">
      <Clock />
      <span className="text-xl font-jamsil">🌦 16°C 1°F</span>
      <span className="text-xl font-jamsil">
        비도 오고 그래서 네 생각이 났어 비도 오고 그래서 네 생각이 났어
      </span>
      <div className="w-3/5	h-36 border-2	border-solid border-black	rounded-xl flex flex-row">
        <div className="w-1/5	h-full border-r-2	border-solid border-black flex flex-col	justify-center items-center text-lg font-light font-jamsil">
          김 재 현
        </div>
        <textarea className="w-full h-full resize-none p-3" />
      </div>
    </div>
  );
};

export default Main;
