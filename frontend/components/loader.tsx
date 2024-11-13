import module from "@/styles/loader.module.css";

const Loader = ({ className }: { className?: string }) => {
  return (
    <div className={className}>
      <div className={module.loader}>
        <div className={`${module.square} ${module.sq1}`}></div>
        <div className={`${module.square} ${module.sq2}`}></div>
        <div className={`${module.square} ${module.sq3}`}></div>
        <div className={`${module.square} ${module.sq4}`}></div>
        <div className={`${module.square} ${module.sq5}`}></div>
        <div className={`${module.square} ${module.sq6}`}></div>
        <div className={`${module.square} ${module.sq7}`}></div>
        <div className={`${module.square} ${module.sq8}`}></div>
        <div className={`${module.square} ${module.sq9}`}></div>
      </div>
    </div>
  );
};

export default Loader;
