package lockfile

func KnownEcosystems() []Ecosystem {
	return []Ecosystem{
		NpmEcosystem,
		NuGetEcosystem,
		CargoEcosystem,
		BundlerEcosystem,
		ComposerEcosystem,
		GoEcosystem,
		MixEcosystem,
		MavenEcosystem,
		PipEcosystem,
		PubEcosystem,
		ConanEcosystem,
		// Disabled temporarily,
		// see https://github.com/threatcode/osv-scanner/pull/128 discussion for additional context
		// AlpineEcosystem,
	}
}
