#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

CORELIB_PKG := go-corelibs/runes
VERSION_TAGS += MAIN
MAIN_MK_SUMMARY := ${CORELIB_PKG}
MAIN_MK_VERSION := v1.0.1

DEPS += golang.org/x/perf/cmd/benchstat

STATS_BENCH       := testdata/bench
STATS_FILE        := ${STATS_BENCH}/${MAIN_MK_VERSION}
STATS_PATH        := ${STATS_BENCH}/${MAIN_MK_VERSION}-d
STATS_FILE_OUTPUT := ${STATS_BENCH}/${MAIN_MK_VERSION}-d/output
STATS_FILE_BYTES  := ${STATS_BENCH}/${MAIN_MK_VERSION}-d/bytes
STATS_FILE_STRING := ${STATS_BENCH}/${MAIN_MK_VERSION}-d/string
STATS_FILE_RUNES  := ${STATS_BENCH}/${MAIN_MK_VERSION}-d/runes

STATS_FILES += ${STATS_FILE}
STATS_FILES += ${STATS_FILE_OUTPUT}
STATS_FILES += ${STATS_FILE_BYTES}
STATS_FILES += ${STATS_FILE_STRING}
STATS_FILES += ${STATS_FILE_RUNES}

.PHONY += benchmark
.PHONY += benchstats

include CoreLibs.mk

benchmark: export BENCH_COUNT=1000
benchmark:
	@rm -fv    "${STATS_FILE}" || true
	@rm -rfv   "${STATS_PATH}" || true
	@mkdir -vp "${STATS_PATH}"
	@$(MAKE) bench | egrep -v '^make' > "${STATS_FILE_OUTPUT}"
	@cat "${STATS_FILE_OUTPUT}" > "${STATS_FILE}"
	@cat "${STATS_FILE_OUTPUT}" \
		| egrep -v '(String|Runes)Reader' \
		| perl -pe 's!(Bytes|String|Runes)Reader!!' > ${STATS_FILE_BYTES}
	@cat "${STATS_FILE_OUTPUT}" \
		| egrep -v '(Bytes|Runes)Reader' \
		| perl -pe 's!(Bytes|String|Runes)Reader!!' > ${STATS_FILE_STRING}
	@cat "${STATS_FILE_OUTPUT}" \
		| egrep -v '(Bytes|String)Reader' \
		| perl -pe 's!(Bytes|String|Runes)Reader!!' > ${STATS_FILE_RUNES}
	@shasum ${STATS_FILES}

benchstats:
	@pushd ${STATS_PATH} > /dev/null \
		&& ${CMD} benchstat -alpha 0.000005 \
			bytes string runes \
		&& popd > /dev/null
