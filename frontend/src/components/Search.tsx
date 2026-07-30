export default function Search () {
    return (
        <div className="flex items-center bg-slate-900 border border-slate-700 rounded-lg overflow-hidden w-full max-w-md">
  <input
    type="text"
    placeholder="Search courses..."
    className="w-full bg-transparent p-2 text-white outline-none"
  />
  <button className="px-4 text-slate-400 hover:text-emerald-400">
    🔍
  </button>
</div>
    )
}